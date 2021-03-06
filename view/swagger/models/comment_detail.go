// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CommentDetail 简要评论结构
// swagger:model comment_detail
type CommentDetail struct {

	// 子级别评论，最多只有两层，不会有更多层
	// Required: true
	Comments []*CommentSummary `json:"comments"`

	// 评论全部内容
	// Required: true
	Content *string `json:"content"`

	// 创建时间
	// Required: true
	CreatedAt *string `json:"created_at"`

	// 评论唯一标示
	// Required: true
	ID *int64 `json:"id"`

	// 点赞数
	// Required: true
	LikeCount *int64 `json:"like_count"`

	// 父级别评论的id
	// Required: true
	ParentID *int64 `json:"parent_id"`

	// user
	// Required: true
	User *UserSummary `json:"user"`
}

// Validate validates this comment detail
func (m *CommentDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLikeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommentDetail) validateComments(formats strfmt.Registry) error {

	if err := validate.Required("comments", "body", m.Comments); err != nil {
		return err
	}

	for i := 0; i < len(m.Comments); i++ {
		if swag.IsZero(m.Comments[i]) { // not required
			continue
		}

		if m.Comments[i] != nil {
			if err := m.Comments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommentDetail) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *CommentDetail) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *CommentDetail) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CommentDetail) validateLikeCount(formats strfmt.Registry) error {

	if err := validate.Required("like_count", "body", m.LikeCount); err != nil {
		return err
	}

	return nil
}

func (m *CommentDetail) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parent_id", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

func (m *CommentDetail) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommentDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentDetail) UnmarshalBinary(b []byte) error {
	var res CommentDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
