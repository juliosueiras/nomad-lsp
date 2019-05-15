package nomadstructs

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

//"all_at_once",
//"constraint",
//"affinity",
//"spread",
//"datacenters",
//"group",
//"id",
//"meta",
//"migrate",
//"name",
//"namespace",
//"parameterized",
//"periodic",
//"priority",
//"region",
//"reschedule",
//"task",
//"type",
//"update",
//"vault",
//"vault_token",

var NomadSpec = &hcldec.BlockMapSpec{
	LabelNames: []string{
		"job",
	},
	TypeName: "job",
	Nested: &hcldec.ObjectSpec{
		"all_at_once": &hcldec.AttrSpec{
			Name: "all_at_once",
			Type: cty.Bool,
		},
		"constraint": ConstraintSpec,
		"group":      GroupSpec,
	},
}

var ConstraintSpec = &hcldec.BlockSpec{
	TypeName: "constraint",
	Nested: &hcldec.ObjectSpec{
		"attribute": &hcldec.AttrSpec{
			Name: "attribute",
			Type: cty.String,
		},
		"operator": &hcldec.AttrSpec{
			Name: "operator",
			Type: cty.String,
		},
		"value": &hcldec.AttrSpec{
			Name: "value",
			Type: cty.String,
		},
	},
}

var GroupSpec = &hcldec.BlockMapSpec{
	LabelNames: []string{
		"group",
	},
	TypeName: "group",
	Nested: &hcldec.ObjectSpec{
		"constraint": ConstraintSpec,
		"task":       TaskSpec,
	},
}

var TaskSpec = &hcldec.BlockMapSpec{
	LabelNames: []string{
		"task",
	},
	TypeName: "task",
	Nested: &hcldec.ObjectSpec{
		"constraint": ConstraintSpec,
	},
}
