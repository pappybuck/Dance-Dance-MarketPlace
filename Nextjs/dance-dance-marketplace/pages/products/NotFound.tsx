import Link from "next/link"

function notFound() {
    return (
        <>
            <Link href="/">Go Back</Link>
            <h1>404 - Product Not Found</h1>
        </>
    )
}

export default notFound