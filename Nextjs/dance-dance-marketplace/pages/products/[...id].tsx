import { GetStaticPropsContext } from "next";
import Navbar from "../../components/Navbar";
import { Product, Review } from "../../generated/graphql";
import { ssrClient } from "../../lib/graphql";  
import { useSelector } from 'react-redux';
import { selectAuthState } from "../../Redux/Slices/authSlice";
import Link from "next/link";
import ReviewField from "../../components/ReviewField";

export async function getStaticPaths() {
    return {
        paths: [],
        fallback: 'blocking'
    }
}

export async function getStaticProps(conext : GetStaticPropsContext<{id: string}>) {
    const id = conext.params?.id
    let product : Product = await ssrClient.request(`
        {
        Product(id: "${id}") {
        name
        description
        price
        quantity
        reviews{
            id
            name
            user{
                firstName
                lastName
                }
            }
        }
    }`
    ).then((data) => data.Product);
    if (product === null) {
        return {
            redirect: {
                destination: '/products/NotFound',
                permanent: false,
            },
        }
    }
    return {
        props: {
            product: product,
            time: new Date().toString()
        }
    }
}


function Product({product, time} : {product: Product, time: string}) {

    const isLoggedIn = useSelector(selectAuthState); 

    return (
        <>
            <Navbar/>
            <h1>Current Time: {time}</h1>
            <h1>Product Name: {product.name}</h1>
            <h3>Product Description: {product.description}</h3>
            <h3>Product Price: {product.price}</h3>
            <h3>Product Quantity: {product.quantity}</h3>
            {product.reviews && <ul>
                <h3>Reviews </h3>
                {!isLoggedIn && 
                    <h4><Link href={"/Login"}>Log in to review product</Link></h4>
                }
                {product.reviews.length < 1 && isLoggedIn && 
                    <h4> Be the first to review! </h4>
                }
                {isLoggedIn &&
                    <ReviewField prod={product}/>
                }
                {product.reviews?.map((review: Review | null) => {
                    return review && (
                        <li key={review.id}>
                            Review: {review.name}
                            {review.user && <div>By: {review.user.firstName} {review.user.lastName}</div>}
                        </li>
                    )
                })}
                </ul>}
        </>
    )
}

export default Product