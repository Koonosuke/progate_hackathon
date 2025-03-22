package validator

import (
	"fmt"
	"reflect"
)


type ValidatorInterface interface {
	Validate(req interface{}) error
}

// バリデーション処理を実行する構造体
type Validator struct{}

// コンストラクタ
func NewValidator() ValidatorInterface {
	return &Validator{}
}

func (v *Validator) ValidateBoardID(boardID string) error {
	if boardID == "" {
		return fmt.Errorf("boardIDは必須です")
	}
	return nil
}

func (v *Validator) ValidateAccessToken(accessToken string) error {
	if accessToken == "" {
		return fmt.Errorf("accessTokenは必須です")
	}
	return nil
}

// リクエストの構造体全体を検証する。
func (v *Validator) Validate(req interface{}) error {
	// req の型情報を取得
	val := reflect.ValueOf(req)

	// リクエストがポインタである場合、ポインタ先を取得
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// リクエストのフィールドを走査
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		// boardID や accessToken であれば個別にバリデーション
		if fieldName == "BoardID" {
			if err := v.ValidateBoardID(field.String()); err != nil {
				return fmt.Errorf("%s: %w", fieldName, err)
			}
		}
		if fieldName == "AccessToken" {
			if err := v.ValidateAccessToken(field.String()); err != nil {
				return fmt.Errorf("%s: %w", fieldName, err)
			}
		}
	}
	return nil
}
