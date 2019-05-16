package nomadstructs

import (
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func GetDriverSpec(driver string) hcldec.Spec {
	switch driver {
	case "docker":
		return dockerSpec
	case "exec":
	case "raw_exec":
		return execSpec
	case "java":
		return javaSpec
	case "qemu":
		return qemuSpec
	case "rkt":
		return rktSpec
	case "lxc":
		return lxcSpec
	case "Singularity":
		return singularitySpec
	}

	return nil
}

func SanitizeDriverConfig(body *hclsyntax.Body, driver string) *hclsyntax.Body {
	switch driver {
	case "docker":
		delete(body.Attributes, "mounts")
		return body
	}

	return body
}

var dockerSpec = hcldec.ObjectSpec{
	"image": &hcldec.AttrSpec{
		Name:     "image",
		Type:     cty.String,
		Required: true,
	},
	"args": &hcldec.AttrSpec{
		Name: "args",
		Type: cty.List(cty.String),
	},
	"auth_soft_fail": &hcldec.AttrSpec{
		Name: "auth_soft_fail",
		Type: cty.Bool,
	},
	"command": &hcldec.AttrSpec{
		Name: "command",
		Type: cty.String,
	},
	"dns_search_domains": &hcldec.AttrSpec{
		Name: "dns_search_domains",
		Type: cty.List(cty.String),
	},
	"dns_options": &hcldec.AttrSpec{
		Name: "dns_options",
		Type: cty.List(cty.String),
	},
	"dns_servers": &hcldec.AttrSpec{
		Name: "dns_servers",
		Type: cty.List(cty.String),
	},
	"entrypoint": &hcldec.AttrSpec{
		Name: "entrypoint",
		Type: cty.List(cty.String),
	},
	"extra_hosts": &hcldec.AttrSpec{
		Name: "extra_hosts",
		Type: cty.List(cty.String),
	},
	"force_pull": &hcldec.AttrSpec{
		Name: "force_pull",
		Type: cty.Bool,
	},
	"hostname": &hcldec.AttrSpec{
		Name: "hostname",
		Type: cty.String,
	},
	"interactive": &hcldec.AttrSpec{
		Name: "interactive",
		Type: cty.Bool,
	},
	"sysctl": &hcldec.BlockAttrsSpec{
		TypeName:    "sysctl",
		ElementType: cty.String,
	},
	"ulimit": &hcldec.BlockAttrsSpec{
		TypeName:    "ulimit",
		ElementType: cty.String,
	},
	"privileged": &hcldec.AttrSpec{
		Name: "privileged",
		Type: cty.Bool,
	},
	"ipc_mode": &hcldec.AttrSpec{
		Name: "ipc_mode",
		Type: cty.String,
	},
	"ipv4_address": &hcldec.AttrSpec{
		Name: "ipv4_address",
		Type: cty.String,
	},
	"ipv6_address": &hcldec.AttrSpec{
		Name: "ipv6_address",
		Type: cty.String,
	},
	"labels": &hcldec.BlockAttrsSpec{
		TypeName:    "labels",
		ElementType: cty.String,
	},
	"load": &hcldec.AttrSpec{
		Name: "load",
		Type: cty.String,
	},
	"mac_address": &hcldec.AttrSpec{
		Name: "mac_address",
		Type: cty.String,
	},
	"network_aliases": &hcldec.AttrSpec{
		Name: "network_aliases",
		Type: cty.List(cty.String),
	},
	"network_mode": &hcldec.AttrSpec{
		Name: "network_mode",
		Type: cty.String,
	},
	"pid_mode": &hcldec.AttrSpec{
		Name: "pid_mode",
		Type: cty.String,
	},
	"port_map": &hcldec.BlockAttrsSpec{
		TypeName:    "port_map",
		ElementType: cty.String,
	},
	"security_opt": &hcldec.AttrSpec{
		Name: "security_opt",
		Type: cty.List(cty.String),
	},
	"shm_size": &hcldec.AttrSpec{
		Name: "shm_size",
		Type: cty.Number,
	},
	"storage_opt": &hcldec.AttrSpec{
		Name: "storage_opt",
		Type: cty.Map(cty.String),
	},
	"SSL": &hcldec.AttrSpec{
		Name: "SSL",
		Type: cty.Bool,
	},
	"tty": &hcldec.AttrSpec{
		Name: "tty",
		Type: cty.Bool,
	},
	"uts_mode": &hcldec.AttrSpec{
		Name: "uts_mode",
		Type: cty.String,
	},
	"userns_mode": &hcldec.AttrSpec{
		Name: "userns_mode",
		Type: cty.String,
	},
	"volumes": &hcldec.AttrSpec{
		Name: "volumes",
		Type: cty.List(cty.String),
	},
	"volume_driver": &hcldec.AttrSpec{
		Name: "volume_driver",
		Type: cty.String,
	},
	"work_dir": &hcldec.AttrSpec{
		Name: "work_dir",
		Type: cty.String,
	},
	"cap_add": &hcldec.AttrSpec{
		Name: "cap_add",
		Type: cty.List(cty.String),
	},
	"cap_drop": &hcldec.AttrSpec{
		Name: "cap_drop",
		Type: cty.List(cty.String),
	},
	"cpu_hard_limit": &hcldec.AttrSpec{
		Name: "cpu_hard_limit",
		Type: cty.Bool,
	},
	"cpu_cfs_period": &hcldec.AttrSpec{
		Name: "cpu_cfs_period",
		Type: cty.Number,
	},
	"advertise_ipv6_address": &hcldec.AttrSpec{
		Name: "advertise_ipv6_address",
		Type: cty.Bool,
	},
	"readonly_rootfs": &hcldec.AttrSpec{
		Name: "readonly_rootfs",
		Type: cty.Bool,
	},
	"pids_limit": &hcldec.AttrSpec{
		Name: "pids_limit",
		Type: cty.Number,
	},
	"logging": &hcldec.BlockSpec{
		TypeName: "logging",
		Nested: &hcldec.ObjectSpec{
			"type": &hcldec.AttrSpec{
				Name: "type",
				Type: cty.String,
			},
			"config": &hcldec.BlockAttrsSpec{
				TypeName:    "config",
				ElementType: cty.String,
			},
		},
	},
	"auth": &hcldec.BlockSpec{
		TypeName: "auth",
		Nested: &hcldec.ObjectSpec{
			"username": &hcldec.AttrSpec{
				Name: "username",
				Type: cty.String,
			},
			"password": &hcldec.AttrSpec{
				Name: "password",
				Type: cty.String,
			},
			"email": &hcldec.AttrSpec{
				Name: "email",
				Type: cty.String,
			},
			"server_address": &hcldec.AttrSpec{
				Name: "server_address",
				Type: cty.String,
			},
		},
	},
	"mounts": &hcldec.AttrSpec{
		Name: "mounts",
		Type: cty.List(cty.DynamicPseudoType),
	},
	"devices": &hcldec.AttrSpec{
		Name: "devices",
		Type: cty.List(cty.DynamicPseudoType),
	},
}

var execSpec = hcldec.ObjectSpec{
	"command": &hcldec.AttrSpec{
		Name:     "command",
		Type:     cty.String,
		Required: true,
	},
	"args": &hcldec.AttrSpec{
		Name: "args",
		Type: cty.List(cty.String),
	},
}

var javaSpec = hcldec.ObjectSpec{
	"class": &hcldec.AttrSpec{
		Name: "class",
		Type: cty.String,
	},
	"args": &hcldec.AttrSpec{
		Name: "args",
		Type: cty.List(cty.String),
	},
	"class_path": &hcldec.AttrSpec{
		Name: "class_path",
		Type: cty.String,
	},
	"jar_path": &hcldec.AttrSpec{
		Name: "jar_path",
		Type: cty.String,
	},
	"jvm_options": &hcldec.AttrSpec{
		Name: "jvm_options",
		Type: cty.List(cty.String),
	},
}

var qemuSpec = hcldec.ObjectSpec{
	"image_path": &hcldec.AttrSpec{
		Name:     "image_path",
		Type:     cty.String,
		Required: true,
	},
	"accelerator": &hcldec.AttrSpec{
		Name: "accelerator",
		Type: cty.String,
	},
	"graceful_shutdown": &hcldec.AttrSpec{
		Name: "graceful_shutdown",
		Type: cty.Bool,
	},
	"port_map": &hcldec.BlockAttrsSpec{
		TypeName:    "port_map",
		ElementType: cty.Number,
	},
}

var rktSpec = hcldec.ObjectSpec{
	"image": &hcldec.AttrSpec{
		Name:     "image",
		Type:     cty.String,
		Required: true,
	},
	"command": &hcldec.AttrSpec{
		Name: "command",
		Type: cty.String,
	},
	"args": &hcldec.AttrSpec{
		Name: "args",
		Type: cty.String,
	},
	"trust_prefix": &hcldec.AttrSpec{
		Name: "trust_prefix",
		Type: cty.String,
	},
	"insecure_options": &hcldec.AttrSpec{
		Name: "insecure_options",
		Type: cty.List(cty.String),
	},
	"dns_servers": &hcldec.AttrSpec{
		Name: "dns_servers",
		Type: cty.List(cty.String),
	},
	"dns_search_domains": &hcldec.AttrSpec{
		Name: "dns_search_domains",
		Type: cty.List(cty.String),
	},
	"net": &hcldec.AttrSpec{
		Name: "net",
		Type: cty.List(cty.String),
	},
	"port_map": &hcldec.BlockAttrsSpec{
		TypeName:    "port_map",
		ElementType: cty.String,
	},
	"debug": &hcldec.AttrSpec{
		Name: "debug",
		Type: cty.Bool,
	},
	"no_overlay": &hcldec.AttrSpec{
		Name: "no_overlay",
		Type: cty.Bool,
	},
	"volumes": &hcldec.AttrSpec{
		Name: "volumes",
		Type: cty.List(cty.String),
	},
	"group": &hcldec.AttrSpec{
		Name: "group",
		Type: cty.String,
	},
}

var lxcSpec = hcldec.ObjectSpec{
	"template": &hcldec.AttrSpec{
		Name:     "template",
		Type:     cty.String,
		Required: true,
	},
	"log_level": &hcldec.AttrSpec{
		Name: "log_level",
		Type: cty.String,
	},
	"verbosity": &hcldec.AttrSpec{
		Name: "verbosity",
		Type: cty.String,
	},
	"volumes": &hcldec.AttrSpec{
		Name: "verbosity",
		Type: cty.List(cty.String),
	},
}

var singularitySpec = hcldec.ObjectSpec{
	"image": &hcldec.AttrSpec{
		Name:     "image",
		Type:     cty.String,
		Required: true,
	},
	"verbose": &hcldec.AttrSpec{
		Name: "verbose",
		Type: cty.String,
	},
	"debug": &hcldec.AttrSpec{
		Name: "debug",
		Type: cty.String,
	},
	"command": &hcldec.AttrSpec{
		Name: "command",
		Type: cty.String,
	},
	"args": &hcldec.AttrSpec{
		Name: "args",
		Type: cty.List(cty.String),
	},
	"binds": &hcldec.AttrSpec{
		Name: "binds",
		Type: cty.List(cty.String),
	},
	"overlay": &hcldec.AttrSpec{
		Name: "overlay",
		Type: cty.List(cty.String),
	},
	"security": &hcldec.AttrSpec{
		Name: "security",
		Type: cty.List(cty.String),
	},
	"contain": &hcldec.AttrSpec{
		Name: "contain",
		Type: cty.String,
	},
	"workdir": &hcldec.AttrSpec{
		Name: "workdir",
		Type: cty.String,
	},
	"pwd": &hcldec.AttrSpec{
		Name: "pwd",
		Type: cty.String,
	},
}
