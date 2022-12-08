

function traversalFolder() {
    for file in `ls $1`
    do
        if [ -d $1"/"$file ]
        then
            traversalFolder $1"/"$file
        else
            if [ "${file##*.}"x = "proto"x ]
            then
                echo $1"/"$file
                protoc --go-grpc_out=. --go_out=. $1"/"$file
            fi
        fi
    done
}

traversalFolder $1