import { NextApiRequest, NextApiResponse } from "next";


export default async function handler(req: NextApiRequest, res : NextApiResponse) {
    if (req.query.secret !== process.env.REVALIDATE_SECRET) {
        return res.status(401).json({message: "Invalid token"});
    }
    try {
        await res.revalidate(`/products/${req.query.id}`);
        return res.json({revalidated: true});
    } catch(err : any) {
        return res.status(500).json({message: err.message});
    }

}