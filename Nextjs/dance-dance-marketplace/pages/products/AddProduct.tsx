import { gql } from "graphql-request";
import { useRouter } from "next/router";
import React, { useCallback, useState } from "react";
import { useMutation } from "urql";
import Navbar from "../../components/Navbar";

const AddProductMutation = gql`
    mutation AddProduct($description: String!, $name: String!, $price: Float!, $quantity: Int!) {
        AddProduct(description: $description, name: $name, price: $price, quantity: $quantity) {
            id
        }
    }
`


export default function AddProduct() {

    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    const [price, setPrice] = useState<number>();
    const [quantity, setQuantity] = useState<number>();
    const router = useRouter();
    const [state, executeMutation] = useMutation(AddProductMutation);
    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const result = executeMutation({ name, description, price, quantity });
        result.then((res) => {
            if (res.error) {
                console.log(res.error);
            } else {
                router.push(`/products/${res.data.AddProduct.id}`);
            }
        });
    }

    return (
        <>
            <Navbar/>
            <form onSubmit={handleSubmit}>
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Name</span>
                    </label>
                    <input type="text" placeholder="Name" className="input input-bordered" value={name} onChange={e => setName(e.target.value)} />
                </div>
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Description</span>
                    </label>
                    <input type="text" placeholder="Description" className="input input-bordered" value={description} onChange={e => setDescription(e.target.value)} />
                </div>
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Price</span>
                    </label>
                    <input type="number" placeholder="Price" className="input input-bordered" value={price} onChange={e => setPrice(parseFloat(e.target.value))} />
                </div>
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Quantity</span>
                    </label>
                    <input type="number" placeholder="Quantity" className="input input-bordered" value={quantity} onChange={e => setQuantity(parseInt(e.target.value))} />
                </div>
                <div className="form-control">
                    <button type="submit" className="btn btn-primary">Submit</button>
                </div>
            </form>
        </>
    )
}