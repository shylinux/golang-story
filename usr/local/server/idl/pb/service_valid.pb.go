package pb

import "shylinux.com/x/golang-story/src/project/server/infrastructure/development/proto"

func (this *ServiceCreateRequest) Validate() error {

	if err := proto.Valid(this, this.MachineID, "machineID", "required"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Mirror, "mirror", "required"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Cmd, "cmd", "required"); err != nil {
		return err
	}

	return nil
}

func (this *ServiceRemoveRequest) Validate() error {

	if err := proto.Valid(this, this.ServiceID, "serviceID", "required"); err != nil {
		return err
	}

	return nil
}

func (this *ServiceInputsRequest) Validate() error {

	if err := proto.Valid(this, this.Key, "key", "required"); err != nil {
		return err
	}

	return nil
}

func (this *ServiceDeployRequest) Validate() error {

	if err := proto.Valid(this, this.ServiceID, "serviceID", "required"); err != nil {
		return err
	}

	return nil
}

func (this *ServiceInfoRequest) Validate() error {

	if err := proto.Valid(this, this.ServiceID, "serviceID", "required"); err != nil {
		return err
	}

	return nil
}

func (this *ServiceListRequest) Validate() error {

	if err := proto.Valid(this, this.Page, "page", "default 1"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Count, "count", "default 10"); err != nil {
		return err
	}

	return nil
}
