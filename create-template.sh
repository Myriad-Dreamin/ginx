object_name=$1
up_camel_object_name=$(sed -r 's/(^|_)(\w)/\U\2/g' <<< "$object_name")
camel_object_name=$(sed -r 's/(_)(\w)/\U\2/g' <<< "$object_name")
model_file=model/db-layer/${object_name}.go
register_file=model/db-layer/register.go
cp model/db-layer/object.go $model_file
sed -i "s/\"object\"/\"${object_name}\"/g" $model_file
sed -i "s/object/${camel_object_name}/g" $model_file
sed -i "s/Object/${up_camel_object_name}/g" $model_file
sed -i "s/Object{}.migrate,/Object{}.migrate,\n\t\t"${up_camel_object_name}"{}.migrate,/g" $register_file