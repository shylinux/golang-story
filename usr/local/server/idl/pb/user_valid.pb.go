package pb

import "shylinux.com/x/golang-story/src/project/server/infrastructure/development/proto"

func (this *UserCreateRequest) Validate() error {

	if err := proto.Valid(this, this.Username, "username", "length >= 6"); err != nil {
		return err
	}

	return nil
}

func (this *UserRemoveRequest) Validate() error {

	if err := proto.Valid(this, this.UserID, "userID", "required"); err != nil {
		return err
	}

	return nil
}

func (this *UserRenameRequest) Validate() error {

	if err := proto.Valid(this, this.UserID, "userID", "required"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Username, "username", "length >= 6"); err != nil {
		return err
	}

	return nil
}

func (this *UserSearchRequest) Validate() error {

	if err := proto.Valid(this, this.Page, "page", "default 1"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Count, "count", "default 10"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Key, "key", "required"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Value, "value", "required"); err != nil {
		return err
	}

	return nil
}

func (this *UserInfoRequest) Validate() error {

	if err := proto.Valid(this, this.UserID, "userID", "required"); err != nil {
		return err
	}

	return nil
}

func (this *UserListRequest) Validate() error {

	if err := proto.Valid(this, this.Page, "page", "default 1"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Count, "count", "default 10"); err != nil {
		return err
	}

	return nil
}
