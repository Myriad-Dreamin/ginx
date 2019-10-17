object_name=$1
placeholder=$2

function print_help() {
    echo "help: create-template.h <object-name> <placeholder-in-routes>"
    echo "e.g.: create-template.h submission sid"
}

if [[ "$object_name" == "" ]]; then
    print_help
    exit 0
fi

if [[ "$placeholder" == "" ]]; then
    print_help
    exit 0
fi

up_camel_object_name=$(sed -r 's/(^|_)(\w)/\U\2/g' <<< "$object_name")
camel_object_name=$(sed -r 's/(_)(\w)/\U\2/g' <<< "$object_name")
model_file=model/db-layer/${object_name}.go
model_entry_file=model/${object_name}.go
service_entry_file=service/${object_name}.go
sp_file=model/sp-layer/${object_name}.go
service_folder=service/${object_name}/

register_file=model/db-layer/register.go
model_provider_file=model/sp-layer/provider.go
service_provider_file=service/provider.go

cp model/db-layer/object.go $model_file
cp model/object.go $model_entry_file
cp model/sp-layer/object.go $sp_file
cp service/object.go $service_entry_file
cp -r service/object $service_folder
sed -i "s/\"object\"/\"${object_name}\"/g" $model_file
sed -i "s/object/${camel_object_name}/g" $model_file
sed -i "s/Object/${up_camel_object_name}/g" $model_file
sed -i "s/\"object\"/\"${object_name}\"/g" $model_entry_file
sed -i "s/object/${camel_object_name}/g" $model_entry_file
sed -i "s/Object/${up_camel_object_name}/g" $model_entry_file
sed -i "s/\"object\"/\"${object_name}\"/g" $sp_file
sed -i "s/object/${camel_object_name}/g" $sp_file
sed -i "s/Object/${up_camel_object_name}/g" $sp_file
sed -i "s/\"object\"/\"${object_name}\"/g" $service_folder*
sed -i "s/object/${camel_object_name}/g" $service_folder*
sed -i "s/Object/${up_camel_object_name}/g" $service_folder*
sed -i "s/oid/${placeholder}/g" $service_folder*
sed -i "s/\"object\"/\"${object_name}\"/g" $service_entry_file
sed -i "s/object/${camel_object_name}/g" $service_entry_file
sed -i "s/Object/${up_camel_object_name}/g" $service_entry_file
sed -i "s/Object{}.migrate,/Object{}.migrate,\n\t\t"${up_camel_object_name}"{}.migrate,/g" $register_file
sed -i "s|switch[ ]ss[ ]:=[ ]database.(type)[ ]{|switch ss := database.(type) {\n\tcase *${up_camel_object_name}DB:\n\t\ts.${camel_object_name}DB = ss|g" $model_provider_file
sed -i "s|switch[ ]ss[ ]:=[ ]service.(type)[ ]{|switch ss := service.(type) {\n\tcase *${up_camel_object_name}Service:\n\t\ts.${camel_object_name}Service = ss|g" $service_provider_file
sed -i "s|objectDB[ ]\*ObjectDB|objectDB *ObjectDB\n\t${camel_object_name}DB *${up_camel_object_name}DB|g" $model_provider_file
sed -i "s|objectService[ ]\*ObjectService|objectService *ObjectService\n\t${camel_object_name}Service *${up_camel_object_name}Service|g" $service_provider_file
sed -i "s|for[ ]_,[ ]dbResult[ ]:=[ ]range[ ]\[\]dbResult{|for _, dbResult := range []dbResult{\n\t\t{\"${camel_object_name}DB\", types.Decay(model.New${up_camel_object_name}DB(srv.Logger, srv.cfg))},|g" server/db.go
sed -i "s|for[ ]_,[ ]serviceResult[ ]:=[ ]range[ ]\[\]serviceResult{|for _, serviceResult := range []serviceResult{\n\t\t{\"${camel_object_name}DB\", types.Decay(service.New${up_camel_object_name}Service(srv.Logger, srv.DatabaseProvider, srv.cfg))},|g" server/service.go
