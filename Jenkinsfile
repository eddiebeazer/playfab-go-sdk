pipeline {
    agent any
    environment {
        PLAYFAB_SECRET     = credentials('playfab-test-secret-key')
        PLAYFAB_TITLE_ID = credentials('playfab-test-title-id')
    }
    tools {
        go "1.18"
        //dependency-check "8.0.1"
    }
    stages {
        stage('Testing') {
            parallel {
                stage('Dependency Check') {
                    steps {
                        dependencyCheck additionalArguments: '', odcInstallation: '8.0.1', stopBuild: true
                        dependencyCheckPublisher failedTotalCritical: 1, failedTotalHigh: 1, unstableTotalLow: 10, unstableTotalMedium: 5
                    }
                }
                // stage('Code Coverage') {
                //     steps {
                //             echo 'Getting modules'
                //             bat 'go get -u -d ./...'

                //             echo 'Code Coverage'
                //             bat 'gocov test ./... | gocov-xml > coverage.xml'

                //             publishCoverage adapters: [cobertura('coverage.xml')]

                //     }
                // }
                // stage('Unit Tests') {
                //     steps {
                //             echo 'Getting modules'
                //             bat 'go get -u -d ./...'

                //             echo 'JUnit Report'
                //             bat 'go test -v 2>&1 ./... | go-junit-report -set-exit-code > report.xml'

                //             junit testResults: 'report.xml', skipPublishingChecks: false

                //     }
                // }
            }
        }
    }
}
