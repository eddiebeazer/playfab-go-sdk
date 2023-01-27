pipeline {
    agent none
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
                    agent {
                        docker {
                            image 'golang:1.18.1-alpine'
                            label 'DockerLinux'
                        }
                    }
                    steps {
                        script {
                            echo 'Getting modules'
                            bat 'go get -u -d ./...'

                            echo 'Installing Test Reporters'
                            bat 'go install github.com/axw/gocov/gocov@latest'
                            bat 'go install github.com/AlekSi/gocov-xml@latest'

                            echo 'Code Coverage'
                            bat 'gocov test ./... | gocov-xml > coverage.xml'
                        }
                    }
                }
                stage('Unit Tests') {
                    agent {
                        docker {
                            image 'golang:1.18.1-alpine'
                            label 'DockerLinux'
                        }
                    }
                    steps {
                        script {
                            echo 'Getting modules'
                            bat 'go get -u -d ./...'

                            echo 'Installing Test Reporters'
                            bat 'go install github.com/jstemmer/go-junit-report/v2@latest'

                            echo 'JUnit Report'
                            bat 'go test -v 2>&1 ./... | go-junit-report -set-exit-code > report.xml'
                        }
                    }
                }
            }
        }
    }
}
