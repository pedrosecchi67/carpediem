BINDIR="$(CURDIR)"

build:
	@go build carpediem.go

setup-path:
	@if [ -f ~/.bashrc ]; then \
		if ! grep -q "$(BINDIR)" ~/.bashrc; then \
			echo '# carpediem poem database' >> ~/.bashrc; \
			echo 'export PATH="$$PATH:$(BINDIR)"' >> ~/.bashrc; \
			echo "Added $(BINDIR) to ~/.bashrc"; \
		else \
			echo "$(BINDIR) already in ~/.bashrc"; \
		fi; \
	fi
	@if [ -f ~/.zshrc ]; then \
		if ! grep -q "$(BINDIR)" ~/.zshrc; then \
			echo '# carpediem poem database' >> ~/.zshrc; \
			echo 'export PATH="$$PATH:$(BINDIR)"' >> ~/.zshrc; \
			echo "Added $(BINDIR) to ~/.zshrc"; \
		else \
			echo "$(BINDIR) already in ~/.zshrc"; \
		fi; \
	fi
	@echo ""
	@echo "Run 'source ~/.bashrc' or 'source ~/.zshrc' to update your current shell"

uninstall:
	@rm carpediem
	@echo "Please clean up your .bashrc/.zshrc manually"

all: build setup-path

.PHONY: build setup-path uninstall all
