pipeline {
    agent any
    environment {
        PLAYFAB_SECRET     = credentials('playfab-test-secret-key')
        PLAYFAB_TITLE_ID = credentials('playfab-test-title-id')
    }
    tools { go '1.18' }
    options {
        parallelsAlwaysFailFast()
    }
    stages {
        stage('Testing') {
            parallel {
                stage('Code Coverage') {
                    steps {
                        script {
                            echo 'Getting modules'
                            bat '%GOROOT%\\bin\\go get -u -d ./...'

                            echo 'Installing Test Reporters'
                            bat '%GOROOT%\\bin\\go install github.com/axw/gocov/gocov@latest'
                            bat '%GOROOT%\\bin\\go install github.com/AlekSi/gocov-xml@latest'

                            echo 'Code Coverage'
                            bat '%GOROOT%\\bin\\gocov test ./... | %GOROOT%\\bin\\gocov-xml > coverage.xml'
                        }
                    }
                }
                stage('Unit Tests') {
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
