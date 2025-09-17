import json

# Load the raw data from your crawler
with open("results/results.json", "r", encoding="utf-8") as f:
    raw_data = json.load(f)

cleaned_data = {}

# Define keywords to identify real Qiskit content
keywords = [
    "Qiskit", "qubit", "circuit", "gate", "measure",
    "QuantumCircuit", "backend", "simulate", "algorithm",
    "VQE", "entanglement", "superposition", "density matrix"
]

for url, pages in raw_data.items():
    for page in pages:
        filtered_content = []
        for line in page.get("slokas", []):
            line = line.strip()  # Remove extra spaces/newlines

            # Skip empty lines or irrelevant messages
            if not line:
                continue
            if ("error while loading" in line or 
                "repository is archived" in line or 
                "please reload" in line):
                continue

            # Keep lines containing keywords or code-related content
            if any(keyword in line for keyword in keywords) or line.startswith(("qc", "from", "import", "#")):
                filtered_content.append(line)

        if filtered_content:
            cleaned_data[url] = {
                "url": page["url"],
                "title": page["title"].strip(),
                "content": filtered_content
            }

# Save the cleaned data to a new JSON file
with open("results/cleaned_qiskit.json", "w", encoding="utf-8") as f:
    json.dump(cleaned_data, f, ensure_ascii=False, indent=2)

print("Cleaning complete! Cleaned data saved to 'results/cleaned_qiskit.json'")
