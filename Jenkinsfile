pipeline {
    agent any

    stages {
        stage('SSH INTO WEB SERVER') {
            steps {
                sshagent(['jenkins-ssh-key']) {
                    sh '''
                    ssh maanvik@192.168.20.2 << 'EOF'
                    cd app/ticketing/pb/
                    ls

                    git pull
                    cd ..
                    ls
                    docker-compose up -d --build
                    docker-compose up -d
                    docker ps 
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
