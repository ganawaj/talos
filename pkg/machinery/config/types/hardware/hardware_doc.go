// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by hack/docgen tool. DO NOT EDIT.

package hardware

import (
	"github.com/siderolabs/talos/pkg/machinery/config/encoder"
)

func (PCIDriverRebindConfigV1Alpha1) Doc() *encoder.Doc {
	doc := &encoder.Doc{
		Type:        "PCIDriverRebindConfig",
		Comments:    [3]string{"" /* encoder.HeadComment */, "PCIDriverRebindConfig allows to configure PCI driver rebinds." /* encoder.LineComment */, "" /* encoder.FootComment */},
		Description: "PCIDriverRebindConfig allows to configure PCI driver rebinds.",
		Fields: []encoder.Doc{
			{},
			{
				Name:        "name",
				Type:        "string",
				Note:        "",
				Description: "PCI device id",
				Comments:    [3]string{"" /* encoder.HeadComment */, "PCI device id" /* encoder.LineComment */, "" /* encoder.FootComment */},
			},
			{
				Name:        "targetDriver",
				Type:        "string",
				Note:        "",
				Description: "Target driver to rebind the PCI device to.",
				Comments:    [3]string{"" /* encoder.HeadComment */, "Target driver to rebind the PCI device to." /* encoder.LineComment */, "" /* encoder.FootComment */},
			},
		},
	}

	doc.AddExample("", examplePCIDriverRebindConfigAlpha1())

	return doc
}

// GetFileDoc returns documentation for the file hardware_doc.go.
func GetFileDoc() *encoder.FileDoc {
	return &encoder.FileDoc{
		Name:        "hardware",
		Description: "Package hardware provides hardware related config documents.\n",
		Structs: []*encoder.Doc{
			PCIDriverRebindConfigV1Alpha1{}.Doc(),
		},
	}
}