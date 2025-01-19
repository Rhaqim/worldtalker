# Use an official Python runtime as a parent image
FROM python:3.11-slim

# Set the working directory in the container
WORKDIR /app

# Install Poetry
RUN pip install poetry

# Disable Poetry's virtual environment creation (use system Python environment)
RUN poetry config virtualenvs.create false

# Copy the pyproject.toml and poetry.lock files to the container
COPY pyproject.toml poetry.lock /app/

# Install the dependencies
RUN poetry install --no-root

# Copy the rest of the application code to the container
COPY . /app

# Ensure all dependencies are in the Python PATH
ENV PYTHONPATH="/app"

# Expose the port that gRPC server will run on
EXPOSE 50051

# Command to run the application
CMD ["poetry", "run", "python", "-m", "wtai.server"]