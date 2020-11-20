#!groovy
pipeline {
    agent {
        label 'jenkins-01'
    }
    parameters {
        choice(name: 'env', choices: ['dev','test','gray','prod'], description: '环境')
        choice(name: 'operator', choices: ['build', 'deploy','restart','rollback'], description: '动作')
        gitParameter branchFilter: 'origin/.*',tagFilter: 'v*', defaultValue: 'master',listSize: '0', name: 'version', type: 'PT_BRANCH_TAG',description: '版本'
    }
    stages {
        // 输出
        stage ("检出代码"){
            steps {
               script {
                  echo "构建环境：${params.env}"
                  echo "构建版本：${params.version}"
                  sh "git add . && git reset --hard && git checkout ${params.version}"
               }
            }
        }
        stage ("代码打包"){
            when {
                expression { params.action == 'deploy' }
            }
            steps {
               script {
                  sh "make build"
               }
            }
        }
        // 镜像推送到harbor
//         stage ("推送镜像"){
//             when {
//                 expression { params.action == 'deploy' }
//             }
//             steps {
//                 script {
//                   sh "make push-image env=${params.env} version=${params.version}"
//                }
//             }
//         }
//         stage ("部署镜像"){
//             when {
//                 expression { params.action == 'deploy' }
//             }
//             steps {
//                 script {
//                     sh "make deploy env=${params.env} version=${params.version}"
//                 }
//             }
//         }
//
//         stage ("回滚"){
//             when {
//                 expression { params.action == 'rollback' }
//             }
//             steps {
//                 script {
//                     sh "make rollback env=${params.env} version=${params.version}"
//                 }
//             }
//         }
//
//         stage ("重启"){
//             when {
//                 expression { params.action == 'restart' }
//             }
//             steps {
//                 script {
//                     sh "make restart env=${params.env} version=${params.version}"
//                 }
//             }
//         }
    }
}