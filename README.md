# AI-based Code Generation for Golang – Generating a Simple Linear Regression Using Various AI Tools

**Author:** Charles Lamb  
**Contact Info:** charlamb@gmail.com  
**GitHub Project Address:** [https://github.com/cglamb/AI_Based_Coding](https://github.com/cglamb/AI_Based_Coding)  
**Git Clone Command:** `git clone https://github.com/cglamb/AI_Based_Coding.git`

## Introduction

This project explores the usefulness of various AI tools in generating Go code, specifically focusing on recreating a simple linear regression model. The AI tools under consideration include OpenAI's ChatGPT 4 and GitHub Copilot, among others. The background for this project is an attempt to replicate the Go-based regression analysis previously conducted in one of my projects: [Linear Regression Benchmarking](https://github.com/cglamb/LinearRegression_Benchmarking).

### Large-Language Model – OpenAI's ChatGPT 4

In assessing ChatGPT 4, the aim was to evaluate its ability to generate Go code for a simple linear regression. The prompt used for this assessment is available at [ChatGPT 4 Prompt](https://chat.openai.com/share/f2d02a7f-2cea-4b5d-a9ca-db5b3b7ad335), with a text version included in the repository as `ChatGPT4_Prompting.txt`.

#### ChatGPT 4 Conclusions

ChatGPT 4 successfully generated code that performed the regression analysis and produced the correct output. However, an issue arose due to a typo in the library name "montanaflynn/stat" instead of the correct "montanaflynn/stats". This error, initially a user mistake, was perpetuated by ChatGPT 4, which claimed to recognize the incorrect library. This problem highlights the importance of prompt engineering, as a second attempt with a modified prompt corrected the mistake, leading ChatGPT 4 to state it did not recognize the misspelled library.

### GitHub Copilot Conclusions

GitHub Copilot's evaluation focused on its auto-completion and AI-based code proposal features, intentionally avoiding CoPilot chat due to its use of an OpenAI LLM, which would overlap with the ChatGPT test. Copilot's recommendations became significantly more accurate after the initial 10-15 lines of code, providing excellent suggestions that improved coding efficiency. Despite this, Copilot's effectiveness still depends on the user's ability to guide the AI, especially in more complex programming scenarios.
