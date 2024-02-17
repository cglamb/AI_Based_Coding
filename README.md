# AI-based Code Generation for Golang – Generating a Simple Linear Regression Using Various AI Tools

**Author:** Charles Lamb  
**Contact Info:** charlamb@gmail.com  
**Github address for this project:** [https://github.com/cglamb/AI_Based_Coding](https://github.com/cglamb/AI_Based_Coding)  
**Git Clone Command for the repository:** `git clone https://github.com/cglamb/AI_Based_Coding.git`

## Introduction

The project explores the usefulness of various AI tools in generating Go code. Various AI/code generation tools are used to recreate a simple linear regression including:

- **OpenAI's ChatGPT 4**
- **Github CoPilot**
- **Jennifer**  

As background, the project attempts to replicate the go-based regression done in an earlier version of one of my projects: [https://github.com/cglamb/LinearRegression_Benchmarking](https://github.com/cglamb/LinearRegression_Benchmarking)

## Recommendation

AI coding tools can be extremely helpful in developing code. The case study above demonstrates how an AI tool could be used to increase efficiency when developing a simple linear regression program in Golang. We would recommend any developer or company employing developers explore the viability of utilizing these tools. While there is likely efficiency advantages in these tools, one clear disadvantage is AI-tools like these will be unfamiliar with custom libraries or code-bases (which may be a significant disadvantage for some developers/companies).  Another disadvantage (as noted in the findings below) is that simple human error can be perpetuated by these tools.  Additionally, it should be noted that this case study was based on a short and fairly common tasks. Follow-up research should be conducted, relating to the effectiveness of these tools in more specialized and complicated applications.  Also recommended is that any user of these tools understand the temperature paramterization of these models, and should consider adjusting temperature downward (if possible).

More info on temperature for Large Lanuage Models - [[https://deepchecks.com/glossary/llm-parameters/]]

## Model Findings and Summary

### ChatGPT 4 Conclusions

**Link to ChatGPT:** [https://chat.openai.com/](https://chat.openai.com/)  
**Prompt available at:** [https://chat.openai.com/share/f2d02a7f-2cea-4b5d-a9ca-db5b3b7ad335](https://chat.openai.com/share/f2d02a7f-2cea-4b5d-a9ca-db5b3b7ad335)  
A txt version of the prompt is included as “ChatGPT4_Prompting.txt” in this repository.

Overall, ChatGPT 4 successfully wrote code that performed the regressions and produced the correct output. Only one issue arose, which was the result of the LLM being asked to use “montanaflynn/stat” instead of “montanaflynn/stats”. The incorrect spelling was perpetuated throughout the code and only resolved once I troubleshooted the issue. While this was primarily a user error, ChatGPT perpetuated the issue saying that it recognized the incorrect library:

“Yes, the montanaflynn/stat library is a statistics library for the Go programming language (often referred to as Golang).”

This issue could have been addressed via better prompt engineering. In a second attempt, ChatGPT 4 was prompted to reduce model temperature. In this new prompt, ChatGPT now said that it did not recognize the incorrectly spelled library. Relevant prompt question and response from the second prompt below:

You:  
We are going to build a regression model in golang? However, before we do that are you familiar with the montanaflynn/stat library for go? If you are not familiar, say you are not familiar. Only answer factually.

ChatGPT:  

I am not familiar with the montanaflynn/stat library for Go.

### GitHub CoPilot Conclusions

**Link to GitHub CoPilot:** [https://github.com/features/copilot](https://github.com/features/copilot)

GitHub Copilot's evaluation focused on its auto-completion and AI-based code proposal features, intentionally avoiding CoPilot chat due to its use of an OpenAI LLM, which would overlap with the ChatGPT test. Copilot's recommendations became significantly more accurate after the initial 10-15 lines of code, providing excellent suggestions that improved coding efficiency. Despite this, Copilot's effectiveness still depends on the user's ability to guide the AI, especially in more complex programming scenarios. In the early phases of testing GitHub CoPilot, I added only minimal commenting. This was a mistake as CoPilot can read comments to contextualize and provide more meaningful suggestions.

### Jennifer Conclusions

**Link to Jennifer:** [https://github.com/dave/jennifer?tab=readme-ov-file](https://github.com/dave/jennifer?tab=readme-ov-file)

Jennifer is a code generator for Go. Jennifer allows for the generation of dynamic or boilerplate code. As the regression analysis performed here is a static code with little repetitive code elements, Jennifer was not used to create the regression code. Instead, Jennifer was used to write unit testing. While only three unit tests were developed, the Jennifer library could be useful in situations with extensive testing.

