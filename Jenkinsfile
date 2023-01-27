pipeline {
    agent {
        label 'DockerLinux'
    }
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
                        }
                    }
                    steps {
                        script {
                            echo 'Getting modules'
                            sh 'go get -u -d ./...'

                            echo 'Installing Test Reporters'
                            sh 'go install github.com/axw/gocov/gocov@latest'
                            sh 'go install github.com/AlekSi/gocov-xml@latest'

                            echo 'Code Coverage'
                            sh 'gocov test ./... | gocov-xml > coverage.xml'
                        }
                    }
                }
                stage('Unit Tests') {
                    agent {
                        docker {
                            image 'golang:1.18.1-alpine'
                        }
                    }
                    steps {
                        script {
                            echo 'Getting modules'
                            sh 'go get -u -d ./...'

                            echo 'Installing Test Reporters'
                            sh 'go install github.com/jstemmer/go-junit-report/v2@latest'

                            echo 'JUnit Report'
                            sh 'go test -v 2>&1 ./... | go-junit-report -set-exit-code > report.xml'
                        }
                    }
                }
            }
        }
    }
}
