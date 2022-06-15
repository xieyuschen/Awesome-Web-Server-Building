# Appendix: some details of Go language
## Embed struct anonymous
As gin Engine embeds a `GroupRouter` anonymously, the engine instance could touch the function relatives to 
the `GroupRouter`.
