import axios from 'axios';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { logout, selectAuthState, selectCurrentEmail, selectCurrentFirstName, selectCurrentLastName, setCredentials } from '../Redux/Slices/authSlice';

export default function Navbar(){
    const email = useSelector(selectCurrentEmail);
    const isLoggedIn = useSelector(selectAuthState);
    const firstName = useSelector(selectCurrentFirstName);
    const lastName = useSelector(selectCurrentLastName);
    const dispatch = useDispatch();
    const router = useRouter();
    useEffect(() => {
        axios.get("https://api.patrickbuck.net/auth/login", {
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
    }, [dispatch]);

    return (
        <div className="navbar bg-primary">
            <div className="flex-1">
                <a className="btn btn-ghost normal-case text-xl">daisyUI</a>
            </div>
            <div className="flex-none gap-2">
                <Link href="/" prefetch={false}>
                    <a className="btn btn-ghost normal-case text-xl">Home</a>
                </Link>
                <Link href="/products/AddProduct" prefetch={false}>
                    <a className="btn btn-ghost normal-case text-xl">Add Product</a>
                </Link>
                <div className="form-control">
                    <input type="text" placeholder="Search" className="input input-bordered" />
                </div>
                    <div className="dropdown dropdown-end">
                <label tabIndex={0} className="btn btn-ghost btn-circle avatar">
                    <div className="w-25 rounded-full bg-yellow-50">
                        Profile
                    </div>
                </label>
                    {isLoggedIn && (
                        <ul tabIndex={0} className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52">
                            <li><a>{email}</a></li>
                            <li><a>{firstName} {lastName}</a></li>
                            <li><a>Profile</a></li>
                            <li><a>Settings</a></li>
                            <li><button onClick={
                                () => {
                                    axios.get("https://api.patrickbuck.net/auth/logout", {
                                        withCredentials: true
                                    }).then((response) => {
                                        dispatch(logout())
                                        router.push("/");
                                    }).catch((error) => {
                                        console.log(error);
                                    });
                                }
                            }>Logout</button></li>
                        </ul>
                    )}
                    {!isLoggedIn && (
                        <ul tabIndex={0} className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52">
                            <li><Link href={'/Login'}>Login</Link></li>
                            <li><Link href='/Register'>Register</Link></li>
                        </ul>
                    )}
                </div>
            </div>
        </div>
    )
}