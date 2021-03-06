// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/grumpylabs/hcf/characteristic"
)

const TypeServiceLabel = "CC"

type ServiceLabel struct {
	*Service

	ServiceLabelNamespace *characteristic.ServiceLabelNamespace

	Name *characteristic.Name
}

func NewServiceLabel() *ServiceLabel {
	svc := ServiceLabel{}
	svc.Service = New(TypeServiceLabel)

	svc.ServiceLabelNamespace = characteristic.NewServiceLabelNamespace()
	svc.AddCharacteristic(svc.ServiceLabelNamespace.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}
