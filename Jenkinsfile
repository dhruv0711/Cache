def username = 'dhruv'


pipeline {
    agent any   
    stages {
        stage('Example stage 1') {
            steps {
                echo 'Hello Mr. ${username}'
            }
        }
        stage('Example stage 2') {
            steps {
               echo "I said, Hello Mr. ${username}"
            }
            stage('Example stage 3') {
            steps {
            echo "Running ${env.BUILD_ID} on ${env.JENKINS_URL}"
            }
        }
    }
}
