def read_content(filename: str) -> str:
    with open(filename) as f:
        template_content = f.read()

    return template_content
