from fastapi import FastAPI
from pydantic import BaseModel
from fastapi.middleware.cors import CORSMiddleware
import requests
import ollama

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class ChatRequest(BaseModel):
    message: str


@app.post("/chat")
async def chat_with_AI(req: ChatRequest):

    response = ollama.chat(
        model='qwen2.5:7b',
        messages=[
            {
                "role": "system",
                "content": "Odpowiadaj krótko i zwięźle po polsku. Maksymalnie 3 zdania."
            },
            {
                "role": "user",
                "content": req.message
            }
        ]
    )

    return {
        "reply": response['message']['content']
    }


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8000)