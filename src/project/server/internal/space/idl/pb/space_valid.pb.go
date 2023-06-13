package pb

import "shylinux.com/x/golang-story/src/project/server/infrastructure/utils/proto"

func (this *SpaceCreateRequest) Validate() error {

	if err := proto.Valid(this, this.Name, "name", "length > 6"); err != nil {
		return err
	}

	return nil
}

func (this *SpaceRemoveRequest) Validate() error {

	if err := proto.Valid(this, this.SpaceID, "spaceID", "required"); err != nil {
		return err
	}

	return nil
}

func (this *SpaceInfoRequest) Validate() error {

	if err := proto.Valid(this, this.SpaceID, "spaceID", "required"); err != nil {
		return err
	}

	return nil
}

func (this *SpaceListRequest) Validate() error {

	if err := proto.Valid(this, this.Page, "page", "default 1"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Count, "count", "default 10"); err != nil {
		return err
	}

	return nil
}
