# faq-chat-bot
Simple chat bot serving answers for frequently asked questions.

**Quick start:**

1. Download application repo.
2. Build an image with app by running `docker build -t faq-chat-bot .` from the repo dir.
3. Start app by `docker-compose up`
4. The chat runs on `localhost:8083` after containers with MySQL and Elastic start properly.

**Create FAQ object:**
```text
POST /faq
    {
        "question":"What is your name?",
        "answer": "My name is Bot. Chat Bot :)"
    }
```

**Create multiple FAQ objects at once:**
```text
POST /faqs
    [
        {
            "question":"What is your name?",
            "answer": "My name is Bot. Chat Bot :)"
        },
        {
            "question":"What is your favourite color?",
            "answer": "My favourite color is blue"
        }
    ]
```
**Update FAQ object by ID:**
```text
PATCH /faq/:id
    {
        "question":"How are you, mr Bot?",
        "answer": "I'm fine."
    }
```

**Get FAQ object by ID:**
```text
GET /faq/:id
```

**Delete FAQ object by ID:**
```text
DELETE /faq/:id
```
