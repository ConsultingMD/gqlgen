// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
"context"
"errors"
"fmt"
"strings"
"github.com/99designs/gqlgen/plugin/federation/fedruntime")







func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.GqlgenService, error) {
	if ec.DisableIntrospection {
		return fedruntime.GqlgenService{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.GqlgenService{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}


func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) ([]fedruntime.fedruntime.GqlgenEntity, error) {
	list := []fedruntime.fedruntime.GqlgenEntity{}
	for _, rep := range representations {
		typeName, ok := rep["__typename"].(string)
		if !ok {
			return nil, errors.New("__typename must be an existing string")
		}
		switch typeName {
		
			
			case "Product":
				id0, err := ec.unmarshalNString2string(ctx, rep["upc"])
					if err != nil {
						return nil, errors.New(fmt.Sprintf("Field %s undefined in schema.", "upc"))
					}
				
	
				entity, err := ec.resolvers.Entity().FindProductByUpc(ctx,
					id0, )
				if err != nil {
					return nil, err
				}
	
				
				list = append(list, entity)
			
		
		default:
			return nil, errors.New("unknown type: "+typeName)
		}
	}
	return list, nil
}

