from festivo import FestivoClient

if __name__ == '__main__':
    c = FestivoClient()
    print(c.get_invoice('inv_123'))
