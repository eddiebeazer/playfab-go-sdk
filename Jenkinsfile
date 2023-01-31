pipeline {
    agent {
        label 'Linux'
    }
    environment {
        PLAYFAB_SECRET     = credentials('playfab-test-secret-key')
        PLAYFAB_TITLE_ID = credentials('playfab-test-title-id')
    }
    tools {
        go '1.18'
    }
    stages {
        stage('Installing Dependencies') {
            steps {
                sh 'go get -u -d ./...'
            }
        }
        stage('Testing') {
            goTest()
        }
    }
}
