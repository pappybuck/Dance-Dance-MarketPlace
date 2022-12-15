import { Product } from "../generated/graphql";


export default function ReviewField({prod} : {prod : Product}){
    return (
        <>
            <h1>Review Go here</h1>
            <textarea></textarea>
        </>
    )
}