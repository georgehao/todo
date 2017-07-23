(function(Vue) {
    "use strict";

    new Vue({
        el: 'body',
        data: {
            loginInfo: {},
            registerModel: {
                email: '',
                password: '',
                confirmPassword: ''
            },
            loginModel: {
                email:'',
                password:''
            }
        },

        methods: {
            register: function () {
                console.log("email:", this.registerModel.email);
                console.log("password:", this.registerModel.password);
                console.log("confirmPassword:", this.registerModel.confirmPassword);
            },
            login: function () {
                console.log("email:", this.loginModel.email);
                console.log("password:", this.loginModel.password);

                this.loginInfo.email = this.loginModel.email;
                this.loginInfo.password = this.loginModel.password;
                Vue.http.options.emulateJSON = true;
                this.$http.post('/login', this.loginInfo).success(function(res) {
                    console.log("login success");
                }).error(function(err) {
                    console.log(err);
                });
            }
        }
    });
})(Vue);