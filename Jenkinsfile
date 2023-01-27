pipeline {
    agent any
    environment {
        PLAYFAB_SECRET     = credentials('playfab-test-secret-key')
        PLAYFAB_TITLE_ID = credentials('playfab-test-title-id')
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

                            publishCoverage adapters: [cobertura("$WORKSPACE\\coverage.xml")], checksName: '', sourceFileResolver: sourceFiles('NEVER_STORE')
                        }
                    }
                }
                stage('Unit Tests') {
                    steps {
                        script {
                            echo 'Getting modules'
                            bat 'go get -u -d ./...'

                            echo 'JUnit Report'
                            bat 'go test -v 2>&1 ./... | E:\\jenkins\\tools\\org.jenkinsci.plugins.golang.GolangInstallation\\1.18\\bin\\bin\\go-junit-report -set-exit-code > report.xml'

                            withChecks('Unit Tests') {
                                junit 'report.xml'
                            }
                        }
                    }
                }
            }
        }
    }
}
