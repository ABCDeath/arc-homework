// auto-generated file, do not edit
package $package

$imports

type $struct_name struct {
	ioc ioc.IoC
	obj object.Object
}

$implementation

func New$struct_name(ioc ioc.IoC, obj object.Object) $interface_name {
	return &$struct_name{
		ioc: ioc,
		obj: obj,
	}
}
