package usecase

import (
	"easylib-go/model"
	"easylib-go/repository"
	"easylib-go/utils"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MembersUsecase interface {
	InsertMember(mmb *model.Member, ctx *gin.Context, req model.MemberCreateRequest) error
	UpdateMember(mmb *model.Member, ctx *gin.Context, req model.MemberCreateRequest) error
	DeleteMember(id string) error
	GetMemberById(id string) (*model.Member, error)
	GetAllMember() ([]*model.Member, error)
}

type memberUsecase struct {
	mmbRepo repository.MemberRepository
}

func (mmbUsecase *memberUsecase) InsertMember(mmb *model.Member, ctx *gin.Context, req model.MemberCreateRequest) error {

	session := sessions.Default(ctx)
	existSession := session.Get("Username")

	for _, member := range req.FormData {

	existPhoneNo, _ := mmbUsecase.mmbRepo.GetMemberByPhoneNumber(mmb.PhoneNo)
	if existPhoneNo != nil {
		return &utils.AppError{
				ErrorCode: 1,
				ErrorMessage: fmt.Sprintf("User data with the PhoneNo %v already exists", existPhoneNo.PhoneNo),
		}
	}
	existIdMember, _ := mmbUsecase.mmbRepo.GetMemberByIdMember(mmb.NoIdentity)
	if existIdMember != nil {
		return &utils.AppError{
			ErrorCode: 1,
				ErrorMessage: fmt.Sprintf("User data with the NoIdentity %v already exists", existIdMember.NoIdentity),
		}
	}
	
	file, _ := member.Open()

	tempFile, err := os.CreateTemp("public", "image-*.jpg")
	if err != nil {
		panic(err)
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
			panic(err)
	}

	tempFile.Write(fileBytes)

	fileName := tempFile.Name()
	newFileName := strings.Split(fileName, "\\")

	CreatedAt := time.Now().UTC()
	UpdatedAt := time.Now().UTC()
	mmb.LoanStatus = false
	mmb.ImagePath = newFileName[1]
	mmb.Id = utils.UuidGenerate()
	mmb.CreatedAt = CreatedAt.Format("2006-01-02 15:04:05")
	mmb.UpdatedAt = UpdatedAt.Format("2006-01-02 15:04:05")
	mmb.CreatedBy = existSession.(string)
	}

	return mmbUsecase.mmbRepo.InsertMember(mmb)
}

func (mmbUsecase *memberUsecase) GetAllMember() ([]*model.Member, error) {
	return mmbUsecase.mmbRepo.GetAllMember()
}

func (mmbUsecase *memberUsecase) UpdateMember(mmb *model.Member, ctx *gin.Context, req model.MemberCreateRequest) error {
	session := sessions.Default(ctx)
	existSession := session.Get("Username")

	for _, member := range req.FormData {

	file, _ := member.Open()

	tempFile, err := os.CreateTemp("public", "image-*.jpg")
	if err != nil {
		panic(err)
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
			panic(err)
	}

	tempFile.Write(fileBytes)

	fileName := tempFile.Name()
	newFileName := strings.Split(fileName, "\\")

	existPhoneNo, _ := mmbUsecase.mmbRepo.GetMemberByPhoneNumber(mmb.PhoneNo)
	if existPhoneNo != nil {
		return &utils.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("User data with the PhoneNo %v already exists", existPhoneNo.PhoneNo),
		}
	}
	existIdMember, _ := mmbUsecase.mmbRepo.GetMemberByIdMember(mmb.NoIdentity)
	if existIdMember != nil {
		return &utils.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("User data with the NoIdentity %v already exists", existIdMember.NoIdentity),
		}
	}

	UpdatedAt := time.Now().UTC()
	mmb.LoanStatus = false
	mmb.ImagePath = newFileName[1]
	mmb.UpdatedAt = UpdatedAt.Format("2006-01-02 15:04:05")
	mmb.CreatedBy = existSession.(string)
	}

	return mmbUsecase.mmbRepo.UpdateMember(mmb)
}

func (mmbUsecase *memberUsecase) GetMemberById(id string) (*model.Member, error) {
	 member, err := mmbUsecase.mmbRepo.GetMemberById(id)
	 if err != nil {
		return nil, fmt.Errorf("error on mmbUsecase.mmbRepo.GetMemberById() : %w ", err)
	 }

	 if member == nil {
		return nil, &utils.AppError{
			ErrorCode: 4,
			ErrorMessage: fmt.Sprintf("Member with id %s not found", id),
		}
	 }

	 return member, nil
}

func (mmbUsecase *memberUsecase) DeleteMember(id string) error {
	member, _ := mmbUsecase.mmbRepo.GetMemberById(id)
	if member == nil {
		return &utils.AppError{
			ErrorCode: 4,
			ErrorMessage: fmt.Sprintf("Member with id %s not found", id),
		}
	}

	mmbUsecase.mmbRepo.DeleteMember(member)

	os.Remove("public/" + member.ImagePath)

	return nil
}

func NewMembersUsecase(mmbRepo repository.MemberRepository) MembersUsecase {
	return &memberUsecase{mmbRepo: mmbRepo}
}
