func (a *$receiver_name) $method_signature {
	resolvedDependency, err := a.ioc.Resolve(ctx, "$module:$dependency", a.obj$resolve_args)
	if err != nil {
		return $return_type{}, err
	}

	resolvedDependencyCast, ok := resolvedDependency.($return_type)
	if !ok {
		return $return_type{}, err
	}

	return resolvedDependencyCast, nil
}
