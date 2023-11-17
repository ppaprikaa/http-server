pipeline {
    agent {
        node {
            label "staging"
        }
    }

    stages {
        stage("Create backup") {
            steps {
                echo " Creating Backup"
                sh """
                if docker image rm -f http-server:backup; then
	                echo "Previous backup deleted"
                fi

                if docker tag http-server:build http-server:backup; then
                    echo "Backup created"
                fi
                """
            }
        }

        stage("Unit-testing") {
            steps {
                echo "Unit testing started"
            }
        }

        stage("Build") {
            steps {
                echo "Building"
                sh """
                if docker image rm -f http-server:build; then
	                echo "Previous build deleted"
                fi

                docker build -t http-server:build . 
                echo "Artifact build completed"
                """
            }
        }

        stage("Deploy build") {
            steps {
                echo "Deploying build, if failed deploy backup"
                sh """
                if docker container rm -f http-server; then
	                echo "Stopped old artifact instance"
                fi

                args=""

                if ! docker run --restart on-failure:5 -d --name http-server -p 5000:80 http-server:build; then
                    docker run --restart on-failure:5 -d --name http-server -p 5000:80 http-server:backup
                fi
                """
            }
        }

    }
}