package utils

import "easylib-go/model"

func ToImageResponse(image model.Images) model.ImageResponse {
	return model.ImageResponse{
		Id:        image.Id,
		Path:      image.Path,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}
}

func ToMemberResponse(member model.Member) model.MemberResponse {
	return model.MemberResponse{
		Id:         member.Id,
		Name:       member.Name,
		PhoneNo:    member.PhoneNo,
		NoIdentity: member.NoIdentity,
		ImagePath:  member.ImagePath,
		LoanStatus: member.LoanStatus,
		CreatedAt:  member.CreatedAt,
		UpdatedAt:  member.UpdatedAt,
		CreatedBy:  member.CreatedBy,
	}
}