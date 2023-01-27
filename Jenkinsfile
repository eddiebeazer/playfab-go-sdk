pipeline {
    agent any
    environment {
        PLAYFAB_SECRET     = credentials('playfab-test-secret-key')
        PLAYFAB_TITLE_ID = credentials('playfab-test-title-id')
    }
    options {
        parallelsAlwaysFailFast()
    }
    stages {
        stage('Testing') {
            parallel {
                stage('Code Coverage') {
                    tools { go '1.18' }
                    steps {
                        script {
                            echo 'Getting modules'
                            bat 'go get -u -d ./...'

                            echo 'Code Coverage'
                            bat 'gocov test ./... | gocov-xml > coverage.xml'
                        }
                    }
                }
                stage('Unit Tests') {
                    steps {
                        script {
                            echo 'Getting modules'
                            bat 'go get -u -d ./...'

                            echo 'JUnit Report'
                            bat 'go test -v 2>&1 ./... | go-junit-report -set-exit-code > report.xml'
                        }
                    }
                }
            }
        }
    }
}
