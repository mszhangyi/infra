package base

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/mszhangyi/infra"
	"github.com/mszhangyi/infra/utils"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	Check(validate)
	return validate
}

func TranState() ut.Translator {
	Check(translator)
	return translator
}

type ValidatorStarter struct {
	infra.BaseStarter
}

func (v *ValidatorStarter) Init() {
	validate = validator.New()
	//创建消息国际化通用翻译器
	cn := zh.New()
	uni := ut.New(cn, cn)
	var found bool
	translator, found = uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Not found translator: zh")
	}
}

func ValidateStruct(s interface{}) (err error) {
	//验证
	var cnErr string
	err = Validate().Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				fmt.Println(err.Translate(translator))
				cnErr += err.Translate(translator) + "\n"
			}
		}
		return errors.New(cnErr)
	}
	return nil
}

func BindValidate(by []byte, data interface{}) (err error) {
	err = utils.JsonApi.Unmarshal(by, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//验证
	err = Validate().Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				fmt.Println(err)
				//log.Error(err.Translate(translator))
			}
		}
		return err
	}
	return nil
}
