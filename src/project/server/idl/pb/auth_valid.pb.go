package pb

import "shylinux.com/x/golang-story/src/project/server/infrastructure/development/proto"

func (this *AuthRegisterRequest) Validate() error {

	if err := proto.Valid(this, this.Username, "username", "length >= 6"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Password, "password", "length >= 6"); err != nil {
		return err
	}

	return nil
}

func (this *AuthLoginRequest) Validate() error {

	if err := proto.Valid(this, this.Username, "username", "length >= 6"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Password, "password", "length >= 6"); err != nil {
		return err
	}

	return nil
}

func (this *AuthLogoutRequest) Validate() error {

	if err := proto.Valid(this, this.Token, "token", "required"); err != nil {
		return err
	}

	return nil
}

func (this *AuthRefreshRequest) Validate() error {

	if err := proto.Valid(this, this.Token, "token", "required"); err != nil {
		return err
	}

	return nil
}

func (this *AuthVerifyRequest) Validate() error {

	if err := proto.Valid(this, this.Token, "token", "required"); err != nil {
		return err
	}

	return nil
}
