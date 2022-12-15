import { useFormik } from "formik";
import { useRouter } from "next/router";
import Navbar from "../components/Navbar";

type FormValues = {
    email: string;
    firstName: string;
    lastName: string;
    username: string;
    password: string;
    confirmPassword: string;
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
    if (!values.firstName) {
        errors.firstName = 'Required';
    }
    if (!values.lastName) {
        errors.lastName = 'Required';
    }
    if (!values.username) {
        errors.username = 'Required';
    }
    if (!values.password) {
        errors.password = 'Required';
    }
    if (values.password.length < 6) {
        errors.password = 'Password must be at least 6 characters';
    }
    if (!/\d/.test(values.password)) {
        errors.password = 'Password must contain at least one number';
    }
    if (!/[A-Z]/.test(values.password)) {
        errors.password = 'Password must contain at least one uppercase letter';
    }
    if (!values.confirmPassword) {
        errors.confirmPassword = 'Required';
    }
    if (values.password !== values.confirmPassword) {
        errors.confirmPassword = 'Passwords do not match';
    }
    return errors;
}



export default function Register() {
    const router = useRouter();
    const formik = useFormik({
        initialValues: {
            email: '',
            firstName: '',
            lastName: '',
            username: '',
            password: '',
            confirmPassword: '',
        },
        validate,
        onSubmit: values => {
            fetch("https://api.patrickbuck.net/auth/register-admin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(values)
            }).then(response => {
                if (response.ok) {
                    router.push("/Login");
                } else {
                    console.log("failed");
                }
            });
        },
    });
    return (
        <div>
            <Navbar/>
            <h1>Register</h1>
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
                        <span className="label-text">First Name</span>
                    </label>
                    <input
                        id="firstName"
                        name="firstName"
                        type="text"
                        placeholder="First Name"
                        className="input input-bordered"
                        onChange={formik.handleChange}
                        value={formik.values.firstName}
                    /> 
                    {formik.errors.firstName ? <div>{formik.errors.firstName}</div> : null}
                </div>
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Last Name</span>
                    </label>
                    <input
                        id="lastName"
                        name="lastName"
                        type="text"
                        placeholder="Last Name"
                        className="input input-bordered"
                        onChange={formik.handleChange}
                        value={formik.values.lastName}
                    />
                    {formik.errors.lastName ? <div>{formik.errors.lastName}</div> : null}
                </div>
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Username</span>
                    </label>
                    <input
                        id="username"
                        name="username"
                        type="text"
                        placeholder="Username"
                        className="input input-bordered"
                        onChange={formik.handleChange}
                        value={formik.values.username}
                    />
                    {formik.errors.username ? <div>{formik.errors.username}</div> : null}
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
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Confirm Password</span>
                    </label>
                    <input
                        id="confirmPassword"
                        name="confirmPassword"
                        type="password"
                        placeholder="Confirm Password"
                        className="input input-bordered"
                        onChange={formik.handleChange}
                        value={formik.values.confirmPassword}
                    />
                    {formik.errors.confirmPassword ? <div>{formik.errors.confirmPassword}</div> : null}
                </div>
                <button className="btn btn-primary" type="submit">Submit</button>
            </form>
        </div>
    );
}