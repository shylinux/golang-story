package pb

import "shylinux.com/x/golang-story/src/project/server/infrastructure/development/proto"

func (this *MachineCreateRequest) Validate() error {

	if err := proto.Valid(this, this.Hostname, "hostname", "required"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Workpath, "workpath", "required"); err != nil {
		return err
	}

	return nil
}

func (this *MachineRemoveRequest) Validate() error {

	if err := proto.Valid(this, this.MachineID, "machineID", "required"); err != nil {
		return err
	}

	return nil
}

func (this *MachineChangeRequest) Validate() error {

	if err := proto.Valid(this, this.MachineID, "machineID", "required"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Status, "status", "required"); err != nil {
		return err
	}

	return nil
}

func (this *MachineInfoRequest) Validate() error {

	if err := proto.Valid(this, this.MachineID, "machineID", "required"); err != nil {
		return err
	}

	return nil
}

func (this *MachineListRequest) Validate() error {

	if err := proto.Valid(this, this.Page, "page", "default 1"); err != nil {
		return err
	}

	if err := proto.Valid(this, this.Count, "count", "default 10"); err != nil {
		return err
	}

	return nil
}
