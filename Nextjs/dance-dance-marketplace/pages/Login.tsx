import axios from "axios";
import { useFormik } from "formik";
import { useDispatch } from "react-redux";
import { useSelector } from "react-redux";
import Navbar from "../components/Navbar";
import { selectAuthState, selectCurrentEmail, setCredentials } from "../Redux/Slices/authSlice";

type FormValues = {
    email: string;
    password: string;
}


type LoginResponse = {
    email: string;
    firstName: string;
    lastName: string;
}

const validate = (values : FormValues) => {
    const errors : Partial<FormValues> = {};
    if (!values.email) {
        errors.email = 'Required';
    } else if (
        !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i.test(values.email)
    ) {
        errors.email = 'Invalid email address';
    }
    if (!values.password) {
        errors.password = 'Required';
    }
    return errors;
}



export default function Login() {
    const email = useSelector(selectCurrentEmail);
    const isLoggedIn = useSelector(selectAuthState);
    const dispatch = useDispatch();
    const formik = useFormik({
        initialValues: {
            email: '',
            password: '',
        },
        validate,
        onSubmit: values => {
            axios.post<LoginResponse>('https://api.patrickbuck.net/auth/login', {
                email: values.email,
                password: values.password
            }, {
                withCredentials: true
            }).then((response) => {
                dispatch(setCredentials({
                    email: response.data.email,
                    firstName: response.data.firstName,
                    lastName: response.data.lastName
                }));
            }).catch((error) => {
                console.log(error);
            });
        },
    });
    return (
        <div>
            <Navbar/>
            <h1>Login</h1>
            <form onSubmit={formik.handleSubmit}>
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Email</span>
                    </label>
                    <input
                        id="email"
                        name="email"
                        type="email"
                        placeholder="Email"
                        className="input input-bordered"
                        onChange={formik.handleChange}
                        value={formik.values.email}
                    />
                    {formik.errors.email ? <div>{formik.errors.email}</div> : null}
                </div>
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Password</span>
                    </label>
                    <input
                        id="password"
                        name="password"
                        type="password"
                        placeholder="Password"
                        className="input input-bordered"
                        onChange={formik.handleChange}
                        value={formik.values.password}
                    />
                    {formik.errors.password ? <div>{formik.errors.password}</div> : null}
                </div>
                <button className="btn btn-primary" type="submit">Submit</button>
            </form>
            {isLoggedIn && 
                <div>
                    <h1>Welcome {email}</h1>
                </div>
            }
        </div>
    );
}