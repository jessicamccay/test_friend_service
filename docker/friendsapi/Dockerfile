FROM docker.uacf.io/go_build/builder_1.6.x:master

# Install/reset dependencies during container build
# NOTE: Tools for managing dependencies are available at https://cgit.uacf.io/go_build.git/plain/tools/.go_service_tools
COPY dependencies.txt dependencies.txt
RUN cached-load-deps.sh dependencies.txt

# Copy sources into the container (see .dockerignore for excluded files)
COPY . .

# Build and test the service app
# Can be scoped to only the friendsapi with this instead:
# RUN ["build.sh", "friends.uacf.io/apps/friendsapi"]
RUN ["build.sh", "friends.uacf.io"]

# Start the service app
CMD ["bin/friends.uacf.io/apps/friendsapi"]

# Expose service app ports
EXPOSE 80 81 82
