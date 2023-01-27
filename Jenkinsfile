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
        stage('Test') {
            agent { label 'DockerLinux' }
            steps {
                script {
                    docker.image('golang:1.18.1-alpine').inside {
                        bat 'go get -u -d ./...'
                    }
                }
            }
        }
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
                //                                /* the return value gets caught and saved into the variable MY_CONTAINER */
                //             MY_CONTAINER = bat(script: '@docker run -d -i python:3.10.7-alpine', returnStdout: true).trim()
                //             echo "mycontainer_id is ${MY_CONTAINER}"
                //    /* python --version gets executed inside the Container */
                //             bat "docker exec ${MY_CONTAINER} python --version "
                //    /* the Container gets removed */
                //             bat "docker rm -f ${MY_CONTAINER}"

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
