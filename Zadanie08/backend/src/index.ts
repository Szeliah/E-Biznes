import express from "express";
import cors from "cors";
import bcrypt from "bcrypt";
import session from "express-session";

import db from "./db.js";

const app = express();

app.use(express.json());
app.use(cors({origin: ["http://localhost:5173", "http://127.0.0.1:5173"], credentials: true, }));
app.use(
    session({
        secret: "super-secret-key",

        resave: false,

        saveUninitialized: false,

        cookie: {
            secure: false,
            maxAge: 1000 * 60 * 60,
        },
    })
);

app.post("/register", async (req, res) => {

    const { email, password } = req.body;

    if (!email || !password) {
        return res.status(400).json({
            message: "Missing data",
        });
    }

    try {

        const existingUser = db.prepare("SELECT * FROM users WHERE email = ?").get(email);

        if (existingUser) {
            return res.status(409).json({
                message: "User already exists",
            });
        }

        const hashedPassword = await bcrypt.hash(password, 10);

        const result = db.prepare(`INSERT INTO users (email, password) VALUES (?, ?)`).run(email, hashedPassword);

        req.session.userId = Number(result.lastInsertRowid);

        res.status(201).json({
            message: "User created",
        });

    } catch (err) {

        console.log(err);

        res.status(500).json({
            message: "Server error",
        });
    }
});

app.post("/login", async (req, res) => {

    const { email, password } = req.body;

    try {

        const user = db.prepare(`
            SELECT * FROM users
            WHERE email = ?
        `).get(email) as
        | {
              id: number;
              email: string;
              password: string;
          }
        | undefined;

        if (!user) {
            return res.status(401).json({
                message: "Invalid credentials",
            });
        }

        const isPasswordCorrect =
            await bcrypt.compare(
                password,
                user.password
            );

        if (!isPasswordCorrect) {
            return res.status(401).json({
                message: "Invalid credentials",
            });
        }

        req.session.userId = user.id;

        res.json({
            message: "Logged in",
        });

    } catch (err) {

        console.error(err);

        res.status(500).json({
            message: "Server error",
        });
    }
});


app.get("/me", (req, res) => {

    if (!req.session.userId) {
        return res.status(401).json({
            message: "Unauthorized",
        });
    }

    const user = db.prepare(`
        SELECT id, email
        FROM users
        WHERE id = ?
    `).get(req.session.userId);

    res.json(user);
});


app.post("/logout", (req, res) => {

    req.session.destroy(() => {

        res.clearCookie("connect.sid");

        res.json({
            message: "Logged out",
        });
    });
});


app.get("/", (req, res) => {
    res.json({
        message: "Backend works",
    });
});


app.listen(3000, () => {
    console.log("Server running");
});