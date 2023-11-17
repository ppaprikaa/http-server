pipeline {
    agent {
        node {
            label "staging"
        }
    }

    stages {
        stage("Create backup") {
            steps {
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

        stage("Build") {
            steps {
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
                sh """
                if docker container rm -f http-server; then
	                echo "Stopped old artifact instance"
                fi

                args="--restart on-failure:5 -d --name http-server -p 5555:5000"

                if ! docker run $args http-server:build; then
                    docker run $args http-server:backup
                fi
                """
            }
        }

    }
}