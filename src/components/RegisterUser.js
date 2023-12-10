import React, {Component} from "react";

class Register extends Component{
    render(){
        return(
            <div>
                <h1>Register Here</h1>
                <form>
                    <div>
                        <label htmlFor="firstName">First Name</label>
                        <input type="text" id="firstName"/>
                    </div>
                    <div>
                        <label htmlFor="lastName">Last Name</label>
                        <input type="text" id="lastName"/>
                    </div>
                    <div>
                        <label htmlFor="age">Age</label>
                        <input type="number" id="age"/>
                    </div>
                    <div>
                        <label htmlFor="userName">Username</label>
                        <input type="text" id="userName"/>
                    </div>
                    <div>
                        <label htmlFor="password">Password</label>
                        <input type="password" id="password"/>
                    </div>

                    <button type="submit">Register</button>
                </form>
            </div>
        )
    }
}

export default Register;