package usecase

import (
	"easylib-go/model"
	"easylib-go/repository"
	"easylib-go/utils"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MembersUsecase interface {
	InsertMember(*gin.Context, *model.Member, model.MemberCreateRequest) []model.MemberResponse
	DeleteMember(string)
	GetMemberById(string) model.MemberResponse
}

type memberUsecase struct {
	mmbRepo repository.MemberRepository
}

func (mmbUsecase *memberUsecase) InsertMember(ctx *gin.Context, mem *model.Member, req model.MemberCreateRequest) []model.MemberResponse {
	session := sessions.Default(ctx)
	existSession := session.Get("Username")
	var imageResponses []model.MemberResponse

	for _, member := range req.FormData { // Menggunakan variable 'member' yang sesuai dengan loop saat ini
		file, _ := member.Open() // Menggunakan 'member.Image' untuk mendapatkan gambar

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
		newMember := model.Member{ // Menggunakan 'newMember' untuk menyimpan data anggota baru
			Id:         utils.UuidGenerate(),
			Name:       mem.Name,
			PhoneNo:    mem.PhoneNo,
			NoIdentity: mem.NoIdentity,
			ImagePath:  newFileName[1],
			LoanStatus: mem.LoanStatus,
			CreatedAt:  CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  UpdatedAt.Format("2006-01-02 15:04:05"),
			CreatedBy:  existSession.(string),
		}

		newMember = mmbUsecase.mmbRepo.InsertMember(newMember) // Memasukkan data anggota baru ke dalam repository
		imageResponses = append(imageResponses, utils.ToMemberResponse(newMember))
	}
	return imageResponses
}


func (mmbUsecase *memberUsecase) DeleteMember(id string) {
	image, err := mmbUsecase.mmbRepo.GetMemberById(id)
	if err != nil {
		panic(err.Error())
	}

	mmbUsecase.mmbRepo.DeleteMember(image)

	os.Remove("public/" + image.ImagePath)
}

func (mmbUsecase *memberUsecase) GetMemberById(id string) model.MemberResponse {
	image, err := mmbUsecase.mmbRepo.GetMemberById(id)
	if err != nil {
		panic(err)
	}
	return utils.ToMemberResponse(image)
}

func NewMembersUsecase(mmbRepo repository.MemberRepository) MembersUsecase {
	return &memberUsecase{mmbRepo: mmbRepo}
}
