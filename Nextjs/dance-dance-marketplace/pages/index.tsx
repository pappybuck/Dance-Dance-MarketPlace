import type { NextPage } from 'next'
import Link from 'next/link'
import Navbar from '../components/Navbar'
import { Product, useFetchProductsQuery } from '../generated/graphql'

const Home: NextPage = () => {
  const [result] = useFetchProductsQuery()

  const {data, fetching, error} = result

  return (
    <div>
      <Navbar/>
      {fetching && (
        <div>Loading...</div>
      )}
      {error && (
        <div>{error.message}</div>
      )}
      {(!data || !data.Products) && !error && !fetching && (
        <div>No data</div>
      )}
      {data && data.Products && (
        <ul>
          {data.Products?.map((product: Product | null) => {
            return product && (
              <li key={product.id}>
                <Link href={`/products/${product.id}`}>
                  {product.name}
                </Link>
              </li>
            )
          }
          )}
        </ul>
      )}
      {/* <Link href="http://prometheus.localhost/">Prometheus Dashboard</Link> */}
    </div>
  )

}

export default Home
