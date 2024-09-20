#!/bin/bash
CUR_DIR="$(pwd)"
CONTAINER_NAME="leo-mysql"
MYSQL_ROOT_PASSWORD="password"
MYSQL_VERSION="latest"
DATA_DIR="/tmp/data"

DB_YAML="$CUR_DIR/config/config.yaml"
DB_USER=$(yq eval '.database.user' $DB_YAML)
DB_PASSWORD=$(yq eval '.database.password' $DB_YAML)
DB_HOST=$(yq eval '.database.host' $DB_YAML)
DB_PORT=$(yq eval '.database.port' $DB_YAML)
DB_NAME=$(yq eval '.database.name' $DB_YAML)

exec_container() {
    if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
        if [ "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
            docker exec -it -e MYSQL_PWD="$MYSQL_ROOT_PASSWORD" $CONTAINER_NAME mysql -u root
        else
            echo "MySQL is not running"
        fi
    else
        echo "MySQL doesn't exist"
    fi   
}

create_db_user() {
    if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
        if [ "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
            docker exec -i -e MYSQL_PWD="$MYSQL_ROOT_PASSWORD" $CONTAINER_NAME mysql -u root <<EOF
CREATE DATABASE IF NOT EXISTS $DB_NAME;
CREATE USER IF NOT EXISTS '$DB_USER'@'%' IDENTIFIED BY '$DB_PASSWORD';
GRANT ALL PRIVILEGES ON $DB_NAME.* TO '$DB_USER'@'%';
FLUSH PRIVILEGES;
EXIT;
EOF
        else
            echo "MySQL doesn't start"
        fi
    else
        echo "MySQL doesn't exist"
    fi   
}

start_container() {
    if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
        if [ "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
            echo "MySQL is already running"
        else
            echo "Starting MySQL..."
            docker start $CONTAINER_NAME
            echo "MySQL started"
        fi
    else
        echo "Creating a new MySQL..."
        mkdir $DATA_DIR
        docker run --name $CONTAINER_NAME \
            -v $DATA_DIR:/var/lib/mysql \
            -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
            -p 3306:3306 \
            -d mysql:$MYSQL_VERSION
        echo "MySQL created"
    fi
}

stop_container() {
    if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
        if [ "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
            echo "stopping MySQL..."
            docker stop $CONTAINER_NAME
            echo "MySQL is stopped"
        else
            echo "MySQL is stopped"
        fi
    else
        echo "$CONTAINER_NAME doesn't exist."
    fi
}

clean_container() {
    stop_container

    if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
        echo "removing MySQL..."
        docker rm $CONTAINER_NAME
        echo "MySQL has been deleted"
    else
        echo "$CONTAINER_NAME doesn't exist"
    fi

    if [ -d "$DATA_DIR" ]; then
        echo "Cleaning MySQL data..."
        sudo rm -rf $DATA_DIR
        echo "MySQL data has been cleared"
    else
        echo "File $DATA_DIR has been cleared"
    fi
}

case "$1" in
    start)
        start_container
        ;;
    stop)
        stop_container
        ;;
    restart)
        stop_container
        start_container
        ;;
    exec)
        exec_container
        ;;
    create)
        create_db_user
        ;;
    clean)
        clean_container
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|exec|create|clean}"
        exit 1
        ;;
esac