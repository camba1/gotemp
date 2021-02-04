# Check if we got a vault token to be able to login
if [ "$#" -ne 1 ]
  then
    echo "No vault token supplied" >&2
    exit 1
fi

./file/scripts/auditsrv.sh $1
./file/scripts/customersrv.sh $1
./file/scripts/productsrv.sh $1
./file/scripts/promotionsrv.sh $1
./file/scripts/usersrv.sh $1