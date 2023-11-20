import { NextApiRequest, NextApiResponse } from "next";
import { createClient } from "redis";

const redis = createClient({
    url: process.env.REDIS,
});

export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse,
) {
    if (req.method !== "GET") return res.status(405).end();

    await redis.connect();
    const { key } = req.query;

    if (typeof key !== "string") return res.status(400).end();

    const ip = await redis.get(key);
    const ok = ip != null && ip !== "" && ip !== undefined;
    // TODO what if key doesn't exist

    void redis.disconnect();
    res.status(200).json({ ip, ok });
}
