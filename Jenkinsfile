pipeline {
    agent any

    stages {
        stage('SSH INTO WEB SERVER') {
            steps {
                sshagent(['jenkins-ssh-key']) {
                    sh '''
                    ssh maanvik@192.168.20.2 << 'EOF'
                    cd app/ticketing/pb/

                    if [ ! -d ".git" ]; then
                        echo '\nCLONING THE REPOSITORY\n'
                        git clone https://github.com/IEP-TFL-TEAM/ticketing-pb.git . || { echo 'FAILED TO CLONE REPOSITORY'; exit 1; }

                        docker build -t ticketing-pb:latest .

                    else
                        echo '\nREPOSITORY ALREADY EXISTS. FETCHING LATEST CHANGES.\n'
                        git fetch origin main

                        LOCAL_COMMIT=$(git rev-parse HEAD)
                        REMOTE_COMMIT=$(git rev-parse origin/main)

                        if [ "$LOCAL_COMMIT" != "$REMOTE_COMMIT" ]; then
                            echo '\nNEW CHANGES DETECTED. PULLING UPDATES.\n'
                            git pull origin main || { echo 'FAILED TO PULL LATEST CHANGES'; exit 1; }

                            echo '\nSTOPPING AND REMOVING THE OLD CONTAINER...\n'
                            docker stop ticketing-pb || true
                            docker rm ticketing-pb || true

                            echo '\nREMOVING THE OLD DOCKER IMAGE...\n'
                            docker rmi ticketing-pb || true

                            echo '\nBUILDING A NEW DOCKER IMAGE...\n'
                            docker build -t ticketing-pb:latest .
                        else
                            echo '\nNO NEW CHANGES DETECTED. SKIPPING BUILD.\n'
                        fi
                    fi

                    echo '\nENSURING THE DOCKER CONTAINER IS RUNNING...\n'
                    docker ps | grep ticketing-pb > /dev/null
                    if [ $? -ne 0 ]; then
                        echo '\nSTARTING THE DOCKER CONTAINER...\n'
                        docker run -d -p 8090:8090 --name ticketing-pb ticketing-pb:latest
                    else
                        echo '\nDOCKER CONTAINER IS ALREADY RUNNING.\n'
                    fi
                    '''
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
        success {
            echo 'PIPELINE COMPLETED SUCCESSFULLY!'
        }
        failure {
            echo 'PIPELINE FAILED.'
        }
    }
}
