import type { NextPage } from 'next'
import { useQuery } from 'urql'

import { Product, Review } from '../../generated/graphql'

const One: NextPage = () => {
    const num = 1
    const [result] = useQuery({
        query: `{
            Product(id: ${num}){
                id
                name
                description
                price
                quantity
                reviews{
                    name
                }
            }
        }`

    })
    const { data, fetching, error } = result
    if (fetching) {
        return <div>Loading...</div>
    }
    if (error) {
        return <div>{error.message}</div>
    }
    if (!data) {
        return <div>No data</div>
    }
    let product : Product = data.Product
    return (
        <div>
            <h1>Loaded</h1>
            <h2>{product.name}</h2>
            <p>{product.description}</p>
            <p>{product.price}</p>
            <p>{product.quantity}</p>
            <ul>
                {product.reviews?.map((review: Review | null) => {
                    return review && (
                        <li key={review.id}>
                            {review.name}
                        </li>
                    )
                }
                )}
            </ul>
        </div>
    )
}

export default One